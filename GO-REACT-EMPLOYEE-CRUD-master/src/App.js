
import React from 'react'
import './App.css';
import '../node_modules/bootstrap/dist/css/bootstrap.css'
import Home from './pages/Home'
import About from './pages/About'
import Contact from './pages/Contact'
import Navbar from './pages/Navbar'
import NotFound from './pages/NotFound'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import AddEmployee from './pages/employee/AddEmployee'
import EditUser from './pages/employee/EditEmployee'
import ViewEmployee from './pages/employee/ViewEmployee'
function App() {
  return (
    <Router>
      <div className="App">
        <Navbar />
        <Switch>
          <Route exact path="/" component={Home}/>
          <Route exact path="/about" component={About} />
          <Route exact path="/contact" component={Contact} />
          <Route exact path="/employee/add" component={AddEmployee} />
          <Route exact path="/employee/edit/:id" component={EditUser} />
          <Route exact path="/employee/:id" component={ViewEmployee} />
          <Route component={NotFound} />
        </Switch>
      </div> 
    </Router>
  );
}

export default App;
