import React from 'react';

class QuestionList extends React.Component {

  render() {
    return (
      <div className='question-list'>
      {sortedQuestions.map(function(question) {
        return (
          <div className='question-item' key={question.id}>
          <div className='question-voted'>{question.votes}</div>
          <button className='question-vote'
            onClick={() => {onUpvote(question.id); }}
            disabled={question.voted} >
            &#9650;
          </button>
          <div className='question-text'>
          <div className='question-body'>{question.body}</div>
          <div className='question-author'>{question.author}</div>
          </div>
          </div>
        )
      })

      }
      </div>
    );
  }

}

export default QuestionList;
