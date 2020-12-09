import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { connect } from "react-redux";
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import { withRouter } from 'react-router-dom'

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
    cursor: 'pointer',
  },
  appbar: {
    backgroundColor: theme.palette.colors.back,
  }
}));

const HeaderComponent = (props) => {
  const classes = useStyles()

  const redirectToTransactions = () => {
    props.history.push('/')
  }

  return (
    <div className={classes.root}>
      <AppBar position="static" className={classes.appbar}>
        <Toolbar>
          <Typography variant="h6" className={classes.title} onClick={redirectToTransactions}>
            BankApp
          </Typography>
          <Button color="inherit" onClick={redirectToTransactions}>Transactions</Button>
        </Toolbar>
      </AppBar>
    </div>
  );
}

const mapStateToProps = (state) => {
  return {};
};

export const Header = withRouter(connect(mapStateToProps, {
})(HeaderComponent));