import urlShortens from '../apis/urlShortens'
import { CREATE_URL } from "./types";

export const createUrl = formValue => async dispatch => {
    try {
        const res = await urlShortens.post('/api/urlShorten', formValue)
        const createUrl = {
            ...res.data, status: res.status
        }
        console.log(res);
        dispatch({ type: CREATE_URL, payload: createUrl });
    } catch (error) {
        const createUrl = {
            status: error.response.status
        }
        dispatch({ type: CREATE_URL, payload: createUrl });
        console.error("createUrl", error);
    }
}