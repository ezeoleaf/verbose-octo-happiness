import React from 'react';
import { HomePage } from './components/Home/HomePage'
import './App.css';
import {
  BrowserRouter,
  Switch,
  Route
} from "react-router-dom"

function App() {
  return (
    <BrowserRouter>
      <Switch>
        <Route exact path="/" component={ HomePage } />
      </Switch>
    </BrowserRouter>
  );
}

export default App;
