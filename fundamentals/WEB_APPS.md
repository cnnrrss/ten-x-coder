# Web Apps

Differentiate between **Isomorphic**, **server-side rendered**, and **single-page apps**:
- **Server-rendered apps** handle each action by making a new request
	- Take advantage of fast perceived performance and SEO-friendly rendering
- **SPAs** handle loading hte content, and responding to user insteractions entirely in browser. 
	- Take advantage of browser push history and XMLHttpRequest
		- Example: Opening a modal
- **Isomorphic web apps** are a combination of these two approaches (like TA was)
Pros:
	- Simplified and improved SEO - bots and crawlers can read all the data on page load.
	- Performance gains in "user-perceived performance"
	- Maintenance gains
	- Improved accessibility beacuse the user can view the app without JS
Cons:
	- Debugging and testing complexity
	- Managing performance on the server
	- Handling the differences between Node.js and browser

The header is static and can be rendered immediately on the server.
		
Products are populated by an API call. On the server, the app fetches this data before rendering the page.

On the initial browser load, the page will already have the data. The browser only has to render the DOM.

If you want to reuse as much code as possible app can rely on Javascript's ability to run in multiple environments (browser, Node.js on server). 

#### Example Stack:
**Node.js**
**Express** - handle routing in server
**React** - HTML components (the view), also Angular, Vue, Ember (olddd), etc.
	- Take advantage of ability to render on the server and build up a complete HTML response to be server to the browser
**React Router** - Build in data prefetching for routes
**Redux** - single-direction data flow**
**Babel**
**webpack** - compile the code that runs in browser, enable nod.js packages in browser


Differentiate the first time the user loads the app from subsequent requests.

**"Bootstrapping"** an application means running the code required to get everything set up. This code is run only one on initial load of application.

**Google PageSpeed Insights** https://developers.google.com/speed/

**React and Redux** mesh well because they both use _"functional programming ideas"_, also community best practice.

**Redux** holds the state of your app in its "store", "actions" (Js objects that represent discrete change of app state) are dispatched from the views. These actions trigger "reducers". Reducers are pure fuctions that take a change and return a new store after responding to the change.
- Key: In Redux, only reducers can update the store. All other components can only _read_ from the store. Additionally, the store is immutable."

**Webpack** - build tool that makes packaging code into a single bundle easy. Has plugin system in loaders allowing access to Babel, Less / Sass / PostCSS compiling.
 - Key to isomorphic. Bundle all dependencies together. Also allows environment variables.  `if (NODE_ENV.IS_BROWSER) { // execute code }`
 