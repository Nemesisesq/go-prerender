import React from 'react';
import {useDispatch, useSelector} from "react-redux";
import {decrement as dec, increment as inc} from "../reducers";

const Counter = props => {

    const count = useSelector(state => {
        debugger
        return state.count;
    }, null);
    const dispatch = useDispatch();

    debugger

    const increment = () => dispatch(inc());
    const decrement = () => dispatch(dec());

    return (
        <div>
            Counter!!!!!!!
            <div>{count}</div>
            <button onClick={increment}>+</button>
            <button onClick={decrement}>-</button>
        </div>
    );
};

export default Counter;