import React from 'react';
import { render } from 'react-dom';

import { Provider } from 'react-redux';

import store from './app/state';
import './index.sass';

import App from './app/App';

const Root = () => (
    <Provider store={store}>
        <App/>
    </Provider>
);

render(<Root />, document.querySelector('#root'));