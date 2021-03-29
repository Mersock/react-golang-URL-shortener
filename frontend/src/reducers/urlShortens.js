import { CREATE_URL, GET_URL } from "../actions/types"


const initialState = {
    createUrl: {
        status: null
    },
    listUrl: []
}

const urlShortens = (state = initialState, { type, payload }) => {
    switch (type) {

        case CREATE_URL:
            return { ...state, createUrl: payload }
        case GET_URL:
            return { ...state, listUrl: payload }
        default:
            return state
    }
}

export default urlShortens
