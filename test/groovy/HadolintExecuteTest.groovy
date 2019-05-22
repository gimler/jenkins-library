import hudson.AbortException

import org.junit.Before
import org.junit.Rule
import org.junit.Test
import org.junit.rules.ExpectedException
import org.junit.rules.RuleChain
import util.*

import static org.junit.Assert.assertThat
import static org.hamcrest.Matchers.*

class HadolintExecuteTest extends BasePiperTest {

    private ExpectedException thrown = new ExpectedException().none()
    private JenkinsShellCallRule shellRule = new JenkinsShellCallRule(this)
    private JenkinsDockerExecuteRule dockerExecuteRule = new JenkinsDockerExecuteRule(this)
    private JenkinsStepRule stepRule = new JenkinsStepRule(this)
    private JenkinsReadYamlRule yamlRule = new JenkinsReadYamlRule(this)
    private JenkinsLoggingRule loggingRule = new JenkinsLoggingRule(this)

    @Rule
    public RuleChain ruleChain = Rules
        .getCommonRules(this)
        .around(thrown)
        .around(yamlRule)
        .around(dockerExecuteRule)
        .around(shellRule)
        .around(stepRule)
        .around(loggingRule)

    @Before
    void init() {
        helper.registerAllowedMethod 'stash', [String, String], { name, includes -> assertThat(name, is('hadolintConfiguration')); assertThat(includes, is('.hadolint.yaml')) }
        helper.registerAllowedMethod 'fileExists', [String], { s -> s == 'Dockerfile' }
        helper.registerAllowedMethod 'checkStyle', [Map], { m -> assertThat(m.pattern, is('hadolint.xml')); return 'checkstyle' }
        helper.registerAllowedMethod 'recordIssues', [Map], { m -> assertThat(m.tools, hasItem('checkstyle')) }
        helper.registerAllowedMethod 'archiveArtifacts', [String], { String p -> assertThat('hadolint.xml', is(p)) }
    }

    @Test
    void testHadolintExecute() {
        stepRule.step.hadolintExecute(script: nullScript, juStabUtils: utils, dockerImage: 'hadolint/hadolint:latest-debian')
        assertThat(dockerExecuteRule.dockerParams.dockerImage, is('hadolint/hadolint:latest-debian'))
        assertThat(loggingRule.log, containsString("Unstash content: buildDescriptor"))
        assertThat(shellRule.shell,
            hasItems(
                "curl -L -o .hadolint.yaml https://github.wdf.sap.corp/raw/SGS/Hadolint-Dockerfile/master/.hadolint.yaml",
                "hadolint Dockerfile -f checkstyle > hadolint.xml || exit 0"
            )
        )
    }

    @Test
    void testNoDockerfile() {
        helper.registerAllowedMethod 'fileExists', [String], { false }
        thrown.expect AbortException
        thrown.expectMessage '[hadolintExecute] Dockerfile is not found.'
        stepRule.step.hadolintExecute(script: nullScript, juStabUtils: utils, dockerImage: 'hadolint/hadolint:latest-debian')
    }
}