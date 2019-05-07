import rootReducer from "./reducers";
import {renderToString} from "react-dom/server";
import {Provider} from "react-redux";
import Counter from "./Counter";
import React from "react";
import {createStore} from "redux";

export const renderShit = () => {
    const store = createStore(rootReducer);

    return {
        html: renderToString(
            <Provider store={store}>
                <Counter/>
            </Provider>
        ),
        state: JSON.stringify(store.getState())
    };
};