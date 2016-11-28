import React from 'react';
import Container from '../../components/container/container'
import Info from '../../components/info'
import {Jumbotron} from 'react-bootstrap';
import FacebookLogin from 'react-facebook-login';

const responseFacebook = (response) => {
    console.log(response);
};

const componentClicked = (event) => {
    console.log(event);
};


const info = () => {
    return (
    <div>
        <Jumbotron style={{backgroundColor: '#f1ee00'}}>
            <div className="container">
                <div className="row">
                    <h1>Welcome to <a href="https://github.com/makalkin/goquest">goquest</a>! Just a
                        playground, nothing more.
                    </h1>

                    <p> Start by signing up &nbsp;
                        <FacebookLogin
                            appId="1179783535435590"
                            autoLoad={true}
                            fields="name,email,picture"
                            onClick={componentClicked}
                            callback={responseFacebook}
                            size="medium"
                        />
                    </p>
                </div>
            </div>
        </Jumbotron>
        <Container>
            <Info
                webpack="Webpack"
                babel="Babel"
                eslint="ESLint"
                react="React"
                reactbootstrap="React Bootstrap"
                rctg="React CSS Transition Group"
                redux="Redux"
                reduxform="Redux Form"
                reduxthunk="Redux Thunk"
                axios="Axios"/>
        </Container>
    </div>
    )
};


export default info;