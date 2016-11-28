import React from 'react';

import Container from "../../components/container/container";
import AppNavbar from "../../components/navbar/appnavbar";
import Info from "../../components/info";
import {Jumbotron} from 'react-bootstrap';

export  default class App extends React.Component {
    static propTypes = {
        children: React.PropTypes.node
    };

    static contextTypes = {
        router: React.PropTypes.object.isRequired
    };

    render() {
        return (
            <div>
                <AppNavbar brand="GOQUEST"/>
                {this.props.children}
            </div>
        );
    }
}