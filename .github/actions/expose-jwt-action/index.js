const core = require('@actions/core');

async function run() {
try {
  // Get aud and request token
  const audience = core.getInput('audience');
  console.log(`audience in javascript ${audience}`);
  const jwt = await core.getIDToken(audience);
  core.setOutput("jwt", jwt);
  core.exportVariable('OUTPUT_JWT', jwt);
} catch (error) {
  core.setFailed(error.message);
}
}

run()
