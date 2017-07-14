import { RECEIVED_DISCUSSIONS, FAILED_DISCUSSIONS, LOADING_DISCUSSIONS} from '../types';
const initialState = {
    events: [],
    errors: [],
    loading: false
}

export default (state=initialState, {type, payload}) => {
    
    switch(type) {
        case LOADING_DISCUSSIONS: return {...state, loading: true}
        case RECEIVED_DISCUSSIONS: return {...state, events: payload, loading: false}
        case FAILED_DISCUSSIONS: return {...state, errors: payload, loading: false}
        default: return {...state};
    }
} 