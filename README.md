# Github user activity

A simple CLI tool for fetching and showing recent Github Activity by using [Github Events API](https://docs.github.com/en/rest/activity/events).

Project link: https://roadmap.sh/projects/github-user-activity

## How to run

Clone the repository: 
```
git clone git@github.com:W0ndermaker/github_user_activity.git
```
Navigate to the project directory: 
```
cd github_user_activity
```
Run the following command to build and run the project:
```
go build -o github-activity
./github-activity <username>
```
Example:
```
./github-activity w0ndermaker
```