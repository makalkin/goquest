import React from 'react';
import Container from '../../components/container/container'
import Info from '../../components/info'
import {Jumbotron, Button} from 'react-bootstrap';
import FacebookLogin from 'react-facebook-login';
import {auth} from '../../api'
import {connect} from 'react-redux';
import {Link} from 'react-router';

const componentClicked = (event) => {
    console.log(event);
};


class AboutContainer extends React.Component {

    responseFacebook = (response) => {
        this.props.dispatch(auth.login(response));
        console.log(response);
    };

    render() {
        return (
            <div>
                <Jumbotron style={{backgroundColor: '#f1ee00'}}>
                    <div className="container">
                        <div className="row">
                            <h1>Welcome to <a href="https://github.com/makalkin/goquest">goquest</a>! Just a
                                playground, nothing more.
                            </h1>
                            {!this.props.auth.isAuthenticated ? (
                                <p> Start by signing up &nbsp;
                                    <FacebookLogin
                                        appId="1179783535435590"
                                        autoLoad={true}
                                        fields="name,email,picture"
                                        onClick={componentClicked}
                                        callback={this.responseFacebook}
                                        size="medium"
                                    />
                                </p>
                            ) : (
                                <p> You are ready &nbsp;
                                    <Link to="/">
                                        <Button bsStyle="primary"  bsSize="large" > Let's GO!</Button>
                                    </Link>
                                </p>
                            )}
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
    }
}

function mapStateToProps(state) {
    return {
        auth: state.auth
    }
}

export default connect(mapStateToProps)(AboutContainer);