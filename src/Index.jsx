import 'whatwg-fetch';
import React from 'react';
import ReactDOM from 'react-dom';

import Form from './Form.jsx';
import QuestionList from './QuestionList.jsx';

const postUrl = '/questions';

function readLocalStorage() {
  const currentVotedQuestions = localStorage.votedQuestions || '[]';
  return JSON.parse(currentVotedQuestions);
}

function updateLocalStorage(questionId) {
  const votedQuestions = readLocalStorage();
  votedQuestions.push(questionId);
  localStorage.votedQuestions = JSON.stringify(votedQuestions);
}

function postQuestion(data) {
  return fetch(postUrl, {
    method: 'post',
    headers: {
      "Content-type": "application/json; charset=UTF-8",
      'x-api-token': 'rDA3kcFNMpQNzkmmDnih'
    },
    body: JSON.stringify(data)
  })
}

function fetchQuestions() {
  return fetch(postUrl, {
    headers: {
      'x-api-token': 'rDA3kcFNMpQNzkmmDnih'
    }
  }).then(function(res) {
    return res.json();
  })
}

function upvote(questionId) {
  return fetch(`${postUrl}/${questionId}/upvote`, {
    method: 'post',
    headers: {
      'x-api-token': 'rDA3kcFNMpQNzkmmDnih'
    }
  });
}

class Index extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      questions: []
    };
    this.onUpvote = this.onUpvote.bind(this);
  }

  componentDidMount() {
    const me = this;
    fetchQuestions().then(function(questions) {
      questions = questions || [];
      const currentVotedQuestions = readLocalStorage();
      questions.map(function(question) {
        if (currentVotedQuestions.includes(question.id)) {
          questions.voted = true;
        }
      })
      me.setState({ questions });
    })
  }

  onUpvote(questionId) {
    const me = this;
    upvote(questionId).then(function() {
      const questions = me.state.questions;
      const questionIndex = questions.findIndex(function(question) {
        return question.id === questionId;
      });
      questions[questionIndex].voted = true;
      questions[questionIndex].votes++;
      me.setState(
        {questions},
        updateLocalStorage(questionId)
      );
    });
  }

  render() {
    return (
      <div>
      <h2>Submit a question</h2>
      <Form onSubmit={postQuestion} />
      <h2>Questions</h2>
      <QuestionList questions={this.state.questions} onUpvote={this.onUpvote} />
      </div>
    );
  }

}

ReactDOM.render(<Index />, document.getElementById('root'));
