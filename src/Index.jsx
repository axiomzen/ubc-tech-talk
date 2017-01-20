import React from 'react';
import ReactDOM from 'react-dom';

import Form from './Form.jsx';

function postQuestion(data) {
  alert('question posted');
}

class Index extends React.Component {

  render() {
    return (
      <div>
      <Form onSubmit={postQuestion} />
      </div>
    );
  }

}

ReactDOM.render(<Index />, document.getElementById('root'));
