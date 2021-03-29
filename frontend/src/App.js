import React from 'react';
import {
  BrowserRouter as Router,
  Route, Switch
} from 'react-router-dom';
import Landing from './components/Landing';
import Admin from './components/Admin'
import { Container } from 'reactstrap';


function App() {
  return (
    <Container>
      <Router >
        <div>
          <Switch>
            <Route path="/adminPage" exact component={Admin} />
            <Route path="/" exact component={Landing} />
          </Switch>
        </div>
      </Router>
    </Container>
  );
}

export default App;
