import urlShortens from '../apis/urlShortens'
import { CREATE_URL, GET_URL } from "./types";

export const createUrl = formValue => async dispatch => {
    try {
        const res = await urlShortens.post('/api/urlShorten', formValue)
        const createUrl = {
            ...res.data, status: res.status
        }
        dispatch({ type: CREATE_URL, payload: createUrl });
    } catch (error) {
        const createUrl = {
            status: error.response.status
        }
        dispatch({ type: CREATE_URL, payload: createUrl });
        console.error("createUrl", error);
    }
}

export const getUrl = () => async dispatch => {
    try {
        const res = await urlShortens.get('/api/urlShorten')
        dispatch({ type: GET_URL, payload: res.data });
    } catch (error) {
        console.error("getUrl", error);
    }
}