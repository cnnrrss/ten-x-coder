# Full Stack MERN App
_M_ is for Mongo
_E_ is for Express
_R_ is for React
_N_ is for Node !

### Instructions
1. `mkdir <app-name>`
2. `npm init OR yarn init`
```
Package name: (<app-name>)
version: (1.0.0)
description: an application to help you snap that habit!
entry point: (server.js)
test command:
git repository: git@github.com:connorvanderhook/<app-name>.git
keywords: mern, habit, app, mindfulness, health, awareness, positive vibes
author: Connor Vanderhook
license: (ISC)
```
3. initialize new repository on the command line
```bash
echo "# <app-name>" >> README.md
git init
git add README.md
git commit -m "first commit"
git remote add origin git@github.com:connorvanderhook/<app-name>.git
git push -u origin master
```
4. Install Express (REST api easy framework)
`npm install express`

5. setup express app 
```javascript
const app = express();
```

6.
```javascript
// Handle Cross-Origin Resource Sharing
// if we try to access our api from a different dormain.
app.use((req, res, next) => {
    res.header("Access-Control-Allow-Origin", "*");
    res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
    next();
});
```

7. Listen!
```javascript
app.listen(port, () => {
    console.log(`Server running on port ${port}`);
});
```
8. Setup app routes using express router
```
GET (/all)
POST (/all)
DELETE (/all/:id)
```
9. Mongoose Schema for interacting with mongodb

```bash
curl --data '{"notes": "my_password"}' http://localhost:4000/apis/agenda \
--header 'Content-Type: application/json' \
--header 'Accept: application/json'
```
