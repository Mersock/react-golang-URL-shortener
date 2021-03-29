import urlShortens from '../apis/urlShortens'
import { CREATE_URL } from "./types";

export const createUrl = formValue => async dispatch => {
    const res = await urlShortens.post('/api/urlShorten', formValue)
    dispatch({ type: CREATE_URL, payload: res.data });
}