import React from 'react'
import {Route, IndexRoute} from 'react-router'
import App from '../containers/App'
import NotFound from '../containers/NotFound'
import Home from '../containers/Home'
// import SignInContainer from '../containers/SignIn';
import About from '../containers/About'

import requireAuthentication from '../containers/AuthenticatedComponent';

export const routes = (
    <div>
        {/*<Route path='/signin' component={SignInContainer} />*/}
        <Route path='/' component={requireAuthentication(App)}>
        {/*<Route path='/' component={App}>*/}
            <IndexRoute component={Home}/>
        </Route>
        <Route path='/about' component={About}/>

        <Route path='*' component={NotFound}/>
    </div>
);
