# Gitlab to YouTrack connector

Configure gitlab webhook to call `{hosted url}/gitlab/issue`.

- The app will create a new issue in the configured YouTrack instance in the configured project.
- Set Title/Summary: `#{Gitlab issue iid} - {Gitlab issue title}`. Example: `#5 - User creation does not work`
- Set Description and append a link to the issue: `Generated for [issue in Gitlab]({Gitlab issue link})`

Environment variables:

```
PORT={Port to expose http server at}
GITLAB_TOKEN={random string for validating gitlab webhooks}
YT_TOKEN={YouTrack access token}
YT_URL={YouTrack url without trailing slash (InCloud: include /youtrack)}
YT_PROJECT_ID={YouTrack project id (example: 0-3)}
```
