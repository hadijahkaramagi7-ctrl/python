const { Octokit } = require("@octokit/rest");
const fs = require("fs");

const octokit = new Octokit({ auth: "YOUR_GITHUB_TOKEN" });

async function uploadFile(owner, repo, path, localPath, message) {
  const content = fs.readFileSync(localPath).toString("base64");

  // Get existing SHA if file exists
  let sha;
  try {
    const { data } = await octokit.repos.getContent({ owner, repo, path });
    sha = data.sha;
  } catch (e) {}

  await octokit.repos.createOrUpdateFileContents({
    owner, repo, path,
    message,
    content,
    sha,
    branch: "main",
  });
  console.log(`✅ Uploaded ${path}`);
}

uploadFile("USERNAME", "REPO", "src/app.js", "./app.js", "Add app.js");
