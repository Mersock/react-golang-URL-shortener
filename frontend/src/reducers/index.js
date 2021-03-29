import { combineReducers } from 'redux';
import authReducer from './authReducer';
import urlShortens from './urlShortens'

export default combineReducers({
    auth: authReducer,
    urlShortens: urlShortens
});