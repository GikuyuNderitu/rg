import React from 'react';
import { render } from 'react-dom';

import { Provider } from 'react-redux';

import './index.sass';

import App from './app/App';

const Root = () => (
    <App title="[){[.Name]}(]"/>
);

render(<Root />, document.querySelector('#root'));