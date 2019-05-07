const INITIAL_STATE = {
    count: 0
};

export const increment = () => ({type: "INC"});
export const decrement = () => ({type: "DEC"});

export default (state = INITIAL_STATE, action) => {
    switch (action.type) {
        case "INC":
            return {...state, count: state.count + 1};
        case "DEC":
            return {...state, count: state.count - 1};
        default:
            return state;
    }
};