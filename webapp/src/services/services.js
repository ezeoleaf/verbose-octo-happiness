import axios from 'axios'
import {API_URL} from '../constants/apiConstants'
import {responseSuccess, responseError} from '../utils/response'

export const fetchBankData = async() => {
  let response = null
  
  await axios.get(API_URL + '/transactions/1') //1 is the user id
    .then((resp) => {
      response = responseSuccess(resp.data)
    })
    .catch((error) => {
      response = responseError('Error');
    })

  return response
}
