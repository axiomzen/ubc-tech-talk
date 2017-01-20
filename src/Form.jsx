import React from 'react';

class Form extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      author: '',
      body: ''
    };

    this.onChangeBody = this.onChangeBody.bind(this);
    this.onChangeAuthor = this.onChangeAuthor.bind(this);
  }

  onChangeAuthor(ev) {
    this.setState({
      author: ev.target.value
    });
  }
  
  onChangeBody(ev) {
    this.setState({
      body: ev.target.value
    });
  }

  render() {
    const { onSubmit } = this.props;

    return (
      <form onSubmit={() => {onSubmit(this.state); }}>
      <textarea
        rows='1'
        onChange={this.onChangeBody}
        value={this.state.body}
        placeholder='Ask anything'
      />
      <textarea
        rows='1'
        onChange={this.onChangeAuthor}
        value={this.state.author}
        placeholder='Your name'
      />
      <input type='submit' disabled={!this.state.body} />
      </form>
    );
  }

}

export default Form;
