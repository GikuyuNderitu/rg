import { combineReducers } from 'redux';
import authReducer from './authReducer';
import discussReducer from './discussReducer';

export default combineReducers({
    authReducer,
    discussReducer
})