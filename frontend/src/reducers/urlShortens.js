import { CREATE_URL } from "../actions/types"


const initialState = {
    createUrl: {
        status: null
    }
}

const urlShortens = (state = initialState, { type, payload }) => {
    switch (type) {

        case CREATE_URL:
            return { ...state, createUrl: payload }

        default:
            return state
    }
}

export default urlShortens
