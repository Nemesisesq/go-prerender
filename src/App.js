import {createStore} from 'redux';
import {Provider} from 'react-redux';
import Counter from './Counter';
import rootReducer from './reducers';
import React from "react";

const load = () => {
    const store = createStore(rootReducer);

    return (
        <Provider store={store}>
            <Counter/>
        </Provider>
    );
};

export default load;
