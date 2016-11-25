import {render} from "react-dom";
import Container from "./components/container/container";
import AppNavbar from "./components/navbar/appnavbar";
import PageHeader from "react-bootstrap/lib/PageHeader";
import Info from "./components/info";

const App = () => (
    <div>
        <AppNavbar brand="React Quick Start" />
        <Container>
            <PageHeader>React Quick Start</PageHeader>
            <p>A ReactJS Quick Start project that supports JavaScript ES7 transpilation to ES5 through Babel, linting with ESLint, and bundling via Webpack.</p>
            <header className="jumbotron" style={{backgroundColor: '#f1ee00'}}>
                <div className="container">
                    <div className="row">
                        <h1>Welcome to <a href="https://github.com/makalkin/goquest">goquest</a>! Just a playground, nothing more.
                        </h1>
                        <p></p>
                    </div>
                </div>
            </header>


            <div className="container">
                <div className="row">
                    <div>
                        <h3>APIs</h3>
                        <ul>
                            <li><a href="quests">Quests</a></li>
                        </ul>
                    </div>
                    <div>

                    </div>
                </div>
            </div>

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
                axios="Axios" />
        </Container>
    </div>
);

render(<App />, document.getElementById("app"));
