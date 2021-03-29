import React from 'react';
import { Router, Route, Switch } from 'react-router-dom';
import Landing from './components/Landing';
import history from './history';
import { Container } from 'reactstrap';


function App() {
  return (
    <Container>
      <Router history={history}>
        <div>
          <Switch>
            <Route path="/" exact component={Landing} />
          </Switch>
        </div>
      </Router>
    </Container>
  );
}

export default App;
