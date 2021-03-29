import { CREATE_URL } from "../actions/types"


const initialState = {
    createUrl: {}
}

const urlShortens = (state = initialState, { type, payload }) => {
    switch (type) {

        case CREATE_URL:
            return { ...state, ...payload }

        default:
            return state
    }
}

export default urlShortens
