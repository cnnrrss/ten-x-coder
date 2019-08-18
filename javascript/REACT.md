# React
https://reactjs.org/docs/glossary.html

## Introducing JSX


## Elements
Unlike browser DOM elements, React elements are plain Objects and cheap to create
	- React DOM takes care of updating the DOM to match the React elements.
	- React elements != Components. Elements are what Components are made of.

`<div id="root"></div>`
Called the "root" DOM node because everythign inside it will be managed by React DOM
	
Applications built with just React usually have a single root DOM node. 

If you are integrating React into an existing app, you may have as many isolated root DOM nodes as you like.

#### Updating the Rendered Element
React elements are *immutable!*
Once you create an element, you can't change its children OR attributes
An element is like a single frame in a movie: it represents the UI at a certain point in time

`
function tick() {
  const element = (
    <div>
      <h1>Hello, world!</h1>
      <h2>It is {new Date().toLocaleTimeString()}.</h2>
    </div>
  );
  ReactDOM.render(element, document.getElementById('root'));
}
setInterval(tick, 1000);
`

It calls ReactDOM.render() every second from a setInterval() callback.

In practice, most React apps only call ReactDOM.render() once. 
In the next sections we will learn how such code gets encapsulated into stateful components.

React DOM compares the element and its children to the previous one, and only applies the DOM updates necessary to bring the DOM to the desired state.

Even though we create an element describing the whole UI tree on every tick, only the text node whose contents has changed gets updated by React DOM.

In our experience, thinking about how the UI should look at any given moment rather than how to change it over time eliminates a whole class of bugs.


## Components and Props
`https://reactjs.org/docs/react-component.html`

Let you split the UI into independent pieces in isolation
Conceptually, the are like Javascript functions
They accept arbitrary inputs (“props”) and return React elements to describe what appears on screen.

`
function Welcome(props) {
  return <h1>Hello, {props.name}</h1>;
}
`
Function is valid React component bcus it accepts a single “props” (properties) object argument with data and returns a React element.

We call such components “function components” because they are literally JavaScript functions.

Can also use ES6 class to define a component
`
class Welcome extends React.Component {
  render() {
    return <h1>Hello, {this.props.name}</h1>;
  }
}
`
Two snippets are equivaleny from React's POV
Classes have some additional features.


Components can refer to other components in their output.
`function App() {
  return (
    <div>
      <Welcome name="Sara" />
      <Welcome name="Cahal" />
      <Welcome name="Edite" />
    </div>
  );
}
`

A good rule of thumb is that if a part of your UI is used several times 
	- Button
	- Panel
	- Select List
	- Container
	- Avatar
Or is complex enough on its own 
	- App
	- FeedStory
	- Comment
it is a good candidate to be a reusable components.

### Props are Read-Only
Declaring component as a function or class, it never modify it's own props.
`
function sum(a, b) {
  return a + b;
}
`
^ called "pure" function because no change inputs.

`
function withdraw(account, amount) {
  account.total -= amount;
}
`
^ impure because `account` is changed

React is pretty flexible but it has a single strict rule:

_All React components must act like pure functions with respect to their props._

## State and Lifecycle
State is similar to props, but it is private and fully controlled by the component.

Function Components don't have state. Use a Class.

You can convert a function component like Clock to a class in five steps:

1. Create an ES6 class, with the same name, that extends React.Component.
2. Add a single empty method to it called render().
3. Move the body of the function into the render() method.
4. Replace props with this.props in the render() body.
5. Delete the remaining empty function declaration.
- Cool Idea?

`
class Clock extends React.Component {
  constructor(props) {
    super(props);
    this.state = {date: new Date()};
  } \
  render() {
    return (
      <div>
        <h1>Hello, world!</h1>
        <h2>It is {this.state.date.toLocaleTimeString()}.</h2>
      </div>
    );
  }
}
`
Note how to pass props to base constructor
`
constructor(props) {
    super(props);
    this.state = {date: new Date()};
}
`

In apps with many components its *very* important to free up resources taken
by components when they are destroyed

We want to set up a timer whenever the Clock is rendered to the DOM for the first time. This is called _mounting_ in React.
`componentDidMount() {}`
	- runs after the component output has been rendered to the DOM

We also want to clear that timer whenever the DOM produced by the Clock is removed. This is called _unmounting_ in React.
`componentWillUnmount() {}`

These methods are called “lifecycle methods”.

this.props is set up by React
this.state has a special meaning
you are free to add additional fields to the class manually 
	(like a timer ID)

Finally, we will implement a method called tick() that the Clock component will run every second.
`
tick() {
    this.setState({
      date: new Date()
    });
  }
 `

It will use this.setState() to schedule updates to the component local state


### Use State Correctly (3 rules)

#### 1. Do Not Modify State Directly
For example, this will not re-render a component:
`// Wrong -> this.state.comment = 'hi';`
`// Good -> this.setState({comment: 'hi'});`
The *only* place where you can assign this.state is the constructor.

#### 2. State Updates May Be Asynchronous
React may batch multiple setState() calls into a single update for performance.

this.props and this.state may be updated asynchronously
you should *not* rely on their values for calculating next state.

For example, this code may fail to update the counter:
`
// Wrong
this.setState({
  counter: this.state.counter + this.props.increment,
});
`
Use form of setState() that accepts a function rather than an object. 
Function will receives the previous state and props at the time of the update.
`
// Correct
this.setState((state, props) => ({
  counter: state.counter + props.increment
}));
`
^ arrow function but works with regular
`
// Correct
this.setState(function(state, props) {
  return {
    counter: state.counter + props.increment
  };
});
`
#### 3. State Updates are Merged
When you call setState(), React merges the object you provide into the current state.

For example, your state may contain several independent variables
`this.state{posts: [], comments:[]} // pseudo`
You can update independently with separate calls
`setState(posts:resp.posts)`
`setState(comments:resp.comments)`

The merging is shallow, so `this.setState({comments})` leaves this.state.posts intact, but completely replaces this.state.comments.

### The Data Flows Down
Neither parent nor child components can know if a certain component is stateful or stateless, and they shouldn’t care whether it is defined as a function or a class.

A component may choose to pass its state down as props to its child components:
`<FormattedDate date={this.state.date} />`

FormattedDate component would receive the date in its props and wouldn’t know whether it came from the Clock’s state, from the Clock’s props, or was typed by hand

This is commonly called a “top-down” or “unidirectional” data flow. 

Any state is always owned by some specific component, and any data or UI derived from that state can only affect components “below” them in the tree.

If you imagine a component tree as a waterfall of props, each component’s state is like an additional water source that joins it at an arbitrary point but also flows down.


## Event Handling
Syntax a bit new in JSX
`<button onclick="activateLasers()">
  Activate Lasers HTML
</button>`
`<button onClick={activateLasers}>
  Activate Lasers JSX
</button>`

To prevent default behavior you must return false
`
function ActionLink() {
  function handleClick(e) {
    e.preventDefault();
    console.log('The link was clicked.');
  }
  return (
    <a href="#" onClick={handleClick}>
      Click me
    </a>
  );
}`

When using React you should generally not need to call addEventListener to add listeners to a DOM element after it is created. Instead, just provide a listener when the element is initially rendered.

When you define a component using an ES6 class, a common pattern is for an event handler to be a method on the class. For example, this Toggle component renders a button that lets the user toggle between “ON” and “OFF” states:
	- (See reactifily/src/components/toggle)

You have to be careful about the meaning of this in JSX callbacks. In JavaScript, class methods are not bound by default. If you forget to bind this.handleClick and pass it to onClick, this will be undefined when the function is actually called.

If calling bind annoys you, there are two ways you can get around this. If you are using the experimental public class fields syntax, you can use class fields to correctly bind callbacks.

### Passing Arguments to Event Handlers
Inside a loop it is common to want to pass an extra parameter to an event handler. For example, if id is the row ID, either of the following would work:
`<button onClick={(e) => this.deleteRow(id, e)}>Delete Row</button>` OR...
`<button onClick={this.deleteRow.bind(this, id)}>Delete Row</button>`
The above two lines are equivalent, and use arrow functions and Function.prototype.bind respectively.

### Conditional Rendering
Yah yah, pretty standard. You can use Javascript operators.
if users show row.
if logged in show log out button.
etc...

## Lists and Keys
https://reactjs.org/docs/lists-and-keys.html

