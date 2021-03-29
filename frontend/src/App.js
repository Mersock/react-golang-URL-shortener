import React from 'react';
import { Router, Route, Switch } from 'react-router-dom';
import Landing from './components/Landing';
import history from './history';

function App() {
  return (
    <div>
      <Router history={history}>
        <div>
          <Switch>
            <Route path="/" exact component={Landing} />
          </Switch>
        </div>
      </Router>
    </div>
  );
}

export default App;
