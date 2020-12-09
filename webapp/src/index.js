import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import * as serviceWorker from './serviceWorker';
import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles';
import { Provider } from "react-redux";
import thunk from "redux-thunk";
import {applyMiddleware, createStore, compose} from "redux";
import combinedReducers from "./store/reducers/reducer";

const outerTheme = createMuiTheme({
  palette: {
    primary: {
      light: '#757ce8',
      main: '#bdbdbd',
      dark: '#263238',
      contrastText: '#fff',
    },
    colors: {
      back: '#455A64',
      font: '#FFFFFF',
      fontGrey: '#bdbdbd',
      icon: '#03A9F4',
    },
  },
});

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const store = createStore(combinedReducers, composeEnhancers(applyMiddleware(thunk)));
const el = document.getElementById('root')

ReactDOM.render(
  (<Provider store={store} >
    <ThemeProvider theme={outerTheme} >
        <App />
    </ThemeProvider>
  </Provider>),
  el
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
