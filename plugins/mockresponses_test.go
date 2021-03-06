package plugins_test

var jiraIssueResponse = `
{
    "expand": "renderedFields,names,schema,operations,editmeta,changelog,versionedRepresentations",
    "id": "10092",
    "self": "http://jira.teamsinspace.com:8080/rest/api/2/issue/10092",
    "key": "TIS-41",
    "fields": {
        "issuetype": {
            "self": "http://jira.teamsinspace.com:8080/rest/api/2/issuetype/3",
            "id": "3",
            "description": "A task that needs to be done.",
            "iconUrl": "http://jira.teamsinspace.com:8080/secure/viewavatar?size=xsmall&avatarId=10618&avatarType=issuetype",
            "name": "Task",
            "subtask": false,
            "avatarId": 10618
        },
        "timespent": 403200,
        "project": {
            "self": "http://jira.teamsinspace.com:8080/rest/api/2/project/10001",
            "id": "10001",
            "key": "TIS",
            "name": "Teams in Space",
            "avatarUrls": {
                "48x48": "http://jira.teamsinspace.com:8080/secure/projectavatar?pid=10001&avatarId=10400",
                "24x24": "http://jira.teamsinspace.com:8080/secure/projectavatar?size=small&pid=10001&avatarId=10400",
                "16x16": "http://jira.teamsinspace.com:8080/secure/projectavatar?size=xsmall&pid=10001&avatarId=10400",
                "32x32": "http://jira.teamsinspace.com:8080/secure/projectavatar?size=medium&pid=10001&avatarId=10400"
            }
        },
        "customfield_11000": "com.atlassian.servicedesk.plugins.approvals.internal.customfield.ApprovalsCFValue@671af69a",
        "fixVersions": [
            {
                "self": "http://jira.teamsinspace.com:8080/rest/api/2/version/10003",
                "id": "10003",
                "description": "Version 3",
                "name": "3.0",
                "archived": false,
                "released": false,
                "releaseDate": "2016-12-23"
            }
        ],
        "customfield_11001": null,
        "aggregatetimespent": 403200,
        "resolution": null,
        "customfield_10500": null,
        "customfield_10501": null,
        "customfield_10700": [],
        "customfield_10503": null,
        "customfield_10900": null,
        "customfield_10506": null,
        "resolutiondate": null,
        "workratio": -1,
        "lastViewed": "2017-09-05T13:35:47.434+0000",
        "watches": {
            "self": "http://jira.teamsinspace.com:8080/rest/api/2/issue/TIS-41/watchers",
            "watchCount": 0,
            "isWatching": false
        },
        "created": "2016-06-24T11:01:32.000+0000",
        "priority": {
            "self": "http://jira.teamsinspace.com:8080/rest/api/2/priority/2",
            "iconUrl": "http://jira.teamsinspace.com:8080/images/icons/priorities/critical.svg",
            "name": "Critical",
            "id": "2"
        },
        "customfield_10300": null,
        "labels": [],
        "customfield_10301": null,
        "customfield_10016": null,
        "customfield_10017": [
            "Large",
            "Support",
            "Team"
        ],
        "timeestimate": 0,
        "aggregatetimeoriginalestimate": null,
        "versions": [
            {
                "self": "http://jira.teamsinspace.com:8080/rest/api/2/version/10004",
                "id": "10004",
                "description": "Version 1.5",
                "name": "1.5",
                "archived": false,
                "released": true,
                "releaseDate": "2016-06-24"
            }
        ],
        "issuelinks": [],
        "assignee": {
            "self": "http://jira.teamsinspace.com:8080/rest/api/2/user?username=admin",
            "name": "admin",
            "key": "admin",
            "emailAddress": "tisadmin@veryrealemail.com",
            "avatarUrls": {
                "48x48": "http://jira.teamsinspace.com:8080/secure/useravatar?ownerId=admin&avatarId=10500",
                "24x24": "http://jira.teamsinspace.com:8080/secure/useravatar?size=small&ownerId=admin&avatarId=10500",
                "16x16": "http://jira.teamsinspace.com:8080/secure/useravatar?size=xsmall&ownerId=admin&avatarId=10500",
                "32x32": "http://jira.teamsinspace.com:8080/secure/useravatar?size=medium&ownerId=admin&avatarId=10500"
            },
            "displayName": "Admin Istrator",
            "active": true,
            "timeZone": "Europe/Copenhagen"
        },
        "updated": "2017-09-06T08:13:46.019+0000",
        "status": {
            "self": "http://jira.teamsinspace.com:8080/rest/api/2/status/3",
            "description": "This issue is being actively worked on at the moment by the assignee.",
            "iconUrl": "http://jira.teamsinspace.com:8080/images/icons/statuses/inprogress.png",
            "name": "open",
            "id": "1",
            "statusCategory": {
                "self": "http://jira.teamsinspace.com:8080/rest/api/2/statuscategory/4",
                "id": 4,
                "key": "indeterminate",
                "colorName": "yellow",
                "name": "In Progress"
            }
        },
        "components": [
            {
                "self": "http://jira.teamsinspace.com:8080/rest/api/2/component/10001",
                "id": "10001",
                "name": "Accomodations",
                "description": "Issues concerning hotels, motels, and other accomodations in space"
            },
            {
                "self": "http://jira.teamsinspace.com:8080/rest/api/2/component/10010",
                "id": "10010",
                "name": "Remote APIs",
                "description": "Remote access to our platform for our travel partners"
            }
        ],
        "timeoriginalestimate": null,
        "description": "Currently LodgingController makes an assumption that all the participants in the group are on the same itinerary. Many of our hotel travel providers limit reservations to 8 people. To accommodate larger groups we need to have multiple hotel/motel reservations on one group reservation. \n\nGroup lodging reservations should be a collection of itineraries with a number of providers that the user selects stored as one data structure.",
        "customfield_10010": null,
        "customfield_10011": "TIS-3",
        "customfield_11100": null,
        "customfield_10012": "98",
        "customfield_11101": null,
        "timetracking": {
            "remainingEstimate": "0h",
            "timeSpent": "112h",
            "remainingEstimateSeconds": 0,
            "timeSpentSeconds": 403200
        }
    }
}
`
var jiraTransitionResponse = `
{
    "expand": "transitions",
    "transitions": [
        {
            "id": "11",
            "name": "Start Progress",
            "to": {
                "self": "http://jira.teamsinspace.com:8080/rest/api/2/status/3",
                "description": "This issue is being actively worked on at the moment by the assignee.",
                "iconUrl": "http://jira.teamsinspace.com:8080/images/icons/statuses/inprogress.png",
                "name": "In Progress",
                "id": "3",
                "statusCategory": {
                    "self": "http://jira.teamsinspace.com:8080/rest/api/2/statuscategory/4",
                    "id": 4,
                    "key": "indeterminate",
                    "colorName": "yellow",
                    "name": "In Progress"
                }
            }
        }
    ]
}
`

var labelResponse = `
[
  {
    "id": 544302811,
    "url": "https://api.gh.com/repos/Praqma/phlow-test/labels/Action%20-%20awaiting%20feed-back",
    "name": "Action - awaiting feed-back",
    "color": "6eb82c",
    "default": false
  },
  {
    "id": 545150499,
    "url": "https://api.gh.com/repos/Praqma/phlow-test/labels/Label1",
    "name": "Label1",
    "color": "ededed",
    "default": false
  },
  {
    "id": 545150500,
    "url": "https://api.gh.com/repos/Praqma/phlow-test/labels/Label2",
    "name": "Label2",
    "color": "ededed",
    "default": false
  },
  {
    "id": 544302897,
    "url": "https://api.gh.com/repos/Praqma/phlow-test/labels/Size%202%20-%20medium",
    "name": "Size 2 - medium",
    "color": "208fe5",
    "default": false
  }
]`

var gitHubIssueResponse = `
{
  "id": 1,
  "url": "https://api.github.com/repos/octocat/Hello-World/issues/1347",
  "repository_url": "https://api.github.com/repos/octocat/Hello-World",
  "labels_url": "https://api.github.com/repos/octocat/Hello-World/issues/1347/labels{/name}",
  "comments_url": "https://api.github.com/repos/octocat/Hello-World/issues/1347/comments",
  "events_url": "https://api.github.com/repos/octocat/Hello-World/issues/1347/events",
  "html_url": "https://github.com/octocat/Hello-World/issues/1347",
  "number": 1347,
  "state": "open",
  "title": "Found a bug",
  "body": "I'm having a problem with this.",
  "user": {
    "login": "octocat",
    "id": 1,
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "gravatar_id": "",
    "url": "https://api.github.com/users/octocat",
    "html_url": "https://github.com/octocat",
    "followers_url": "https://api.github.com/users/octocat/followers",
    "following_url": "https://api.github.com/users/octocat/following{/other_user}",
    "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
    "organizations_url": "https://api.github.com/users/octocat/orgs",
    "repos_url": "https://api.github.com/users/octocat/repos",
    "events_url": "https://api.github.com/users/octocat/events{/privacy}",
    "received_events_url": "https://api.github.com/users/octocat/received_events",
    "type": "User",
    "site_admin": false
  },
  "labels": [
    {
      "id": 208045946,
      "url": "https://api.github.com/repos/octocat/Hello-World/labels/bug",
      "name": "bug",
      "color": "f29513",
      "default": true
    }
  ],
  "assignee": {
    "login": "octocat",
    "id": 1,
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "gravatar_id": "",
    "url": "https://api.github.com/users/octocat",
    "html_url": "https://github.com/octocat",
    "followers_url": "https://api.github.com/users/octocat/followers",
    "following_url": "https://api.github.com/users/octocat/following{/other_user}",
    "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
    "organizations_url": "https://api.github.com/users/octocat/orgs",
    "repos_url": "https://api.github.com/users/octocat/repos",
    "events_url": "https://api.github.com/users/octocat/events{/privacy}",
    "received_events_url": "https://api.github.com/users/octocat/received_events",
    "type": "User",
    "site_admin": false
  },
  "milestone": {
    "url": "https://api.github.com/repos/octocat/Hello-World/milestones/1",
    "html_url": "https://github.com/octocat/Hello-World/milestones/v1.0",
    "labels_url": "https://api.github.com/repos/octocat/Hello-World/milestones/1/labels",
    "id": 1002604,
    "number": 1,
    "state": "open",
    "title": "v1.0",
    "description": "Tracking milestone for version 1.0",
    "creator": {
      "login": "octocat",
      "id": 1,
      "avatar_url": "https://github.com/images/error/octocat_happy.gif",
      "gravatar_id": "",
      "url": "https://api.github.com/users/octocat",
      "html_url": "https://github.com/octocat",
      "followers_url": "https://api.github.com/users/octocat/followers",
      "following_url": "https://api.github.com/users/octocat/following{/other_user}",
      "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
      "organizations_url": "https://api.github.com/users/octocat/orgs",
      "repos_url": "https://api.github.com/users/octocat/repos",
      "events_url": "https://api.github.com/users/octocat/events{/privacy}",
      "received_events_url": "https://api.github.com/users/octocat/received_events",
      "type": "User",
      "site_admin": false
    },
    "open_issues": 4,
    "closed_issues": 8,
    "created_at": "2011-04-10T20:09:31Z",
    "updated_at": "2014-03-03T18:58:10Z",
    "closed_at": "2013-02-12T13:22:01Z",
    "due_on": "2012-10-09T23:39:01Z"
  },
  "locked": false,
  "comments": 0,
  "pull_request": {
    "url": "https://api.github.com/repos/octocat/Hello-World/pulls/1347",
    "html_url": "https://github.com/octocat/Hello-World/pull/1347",
    "diff_url": "https://github.com/octocat/Hello-World/pull/1347.diff",
    "patch_url": "https://github.com/octocat/Hello-World/pull/1347.patch"
  },
  "closed_at": null,
  "created_at": "2011-04-22T13:33:48Z",
  "updated_at": "2011-04-22T13:33:48Z",
  "closed_by": {
    "login": "octocat",
    "id": 1,
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "gravatar_id": "",
    "url": "https://api.github.com/users/octocat",
    "html_url": "https://github.com/octocat",
    "followers_url": "https://api.github.com/users/octocat/followers",
    "following_url": "https://api.github.com/users/octocat/following{/other_user}",
    "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
    "organizations_url": "https://api.github.com/users/octocat/orgs",
    "repos_url": "https://api.github.com/users/octocat/repos",
    "events_url": "https://api.github.com/users/octocat/events{/privacy}",
    "received_events_url": "https://api.github.com/users/octocat/received_events",
    "type": "User",
    "site_admin": false
  }
}
`
