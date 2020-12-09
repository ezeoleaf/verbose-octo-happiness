import React, {useState} from 'react'

import { connect } from "react-redux";
import {withRouter} from 'react-router-dom'
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import { fetchBankData } from '../../services/services';
import { Header } from '../Common/Header';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    padding: 20
  },
  paper: {
    padding: theme.spacing(2),
    textAlign: 'center',
    color: theme.palette.colors.font,
    background: 'none',
    border: 'none',
    boxShadow: 'none'
  },
  fab: {
    color: theme.palette.colors.font,
    backgroundColor: theme.palette.colors.icon,
  },
  negative: {
    color: 'red',
  }
}));

const HomeComponent = (props) => {
    const classes = useStyles();
    const [bankData, setBankData] = useState(false)

    const getData = async () => {
      if (bankData) {
        return;
      }

      let resp = await fetchBankData()
      console.log(resp)

      if (resp.success) {
        setBankData(resp.data)
      }
    }

    const displayDate = (d) => {
      let dt = new Date(d)

      return (dt.getMonth()+1) + '/'
            + dt.getDate() + '/'
            + dt.getFullYear()+ ' '
            + dt.getHours() + ":"
            + dt.getMinutes() + ":" + dt.getSeconds(); //TODO: Use Moment
    } 

    getData()

    return (
      <div>
        <Header />
        <div className={classes.root}>
          <Grid container justify="center" alignItems="center" spacing={3}>
            <Grid item xs={12} sm={6}>
              <Paper className={classes.paper}>
                <Typography className="mb-5" variant="h5" component="h5" id="scroll-header">
                  Transactions
                </Typography>
              </Paper>
            </Grid>
            {bankData && (bankData.transactions.map((val, i) => {
                return (
                  <Grid item xs={12} key={i}>
                    <Paper className={classes.paper}>
                      <Typography className="mb-5" variant="h5" component="h5" id="scroll-header">
                        <div className={val.type === "debit" ? classes.negative : ''}>
                {displayDate(val.date)} | {val.entity} | {val.description} | {val.type === "debit" ? '-' : ''}${val.amount}
                        </div>
                      </Typography>
                    </Paper>
                  </Grid>
                )
              }))
            }
          </Grid>
        </div>
      </div>
    );
}

const mapStateToProps = (state) => {
  return {
    ovewview: state.overview,
  };
};

export const HomePage = withRouter(connect(mapStateToProps, {
})(HomeComponent));
