class App extends React.Component {
    componentDidMount() {

    }

    render() {
        if (this.loggedIn) {
            return (<LoggedIn />);
        } else {
            return (<Home />);
        }
    }
};

class Home extends React.Component {
    render() {
        return (
            <div className="container">
                <div className="col-xs-12 jumbotron text-center">
                    <h1>Hello World</h1>
                    <p>Provide valuable feedback to us</p>
                    <a className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
                </div>
            </div>
        );
    }
};

class LoggedIn extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            products: []
        }
    }

    render() {
        return (
            <div class="col-lg-12">
                <span className="pull-right"><a onClick="{this.logout}"></a></span>
                <h2>Welcme to Hello World</h2>
                <p>
                    Below one'll find the latest games that need feedback.
                    Please provide honest feedback so developers can make the best game
                </p>
                <div className="row">
                    {this.state.products.map(function (product, i) {
                        return <Product key={i} product={product} />
                    })}
                </div>
            </div>
        );
    }
};

class Product extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            voted: null
        }
    }

    upvote() { }

    downvote() { }

    render() {
        return (
            <div className="col-xs-4">
                <div className="panel panel-default">
                    <div className="panel-heading">
                        {this.props.product.Name}
                        <span className="pull-right">{this.state.voted}</span>
                    </div>
                    <div className="panel-body">
                        {this.props.product.Description}
                    </div>
                    <div className="panel-footer">
                        <a onClick={this.upvote} className="btn btn-default">
                            <span className="glyphicon glyphicon-thumbs-up"></span>
                        </a>
                        <a onClick={this.downvote} className="btn btn-default pull-right">
                            <span className="glyphicon glyphicon-thumbs-down"></span>
                        </a>
                    </div>
                </div>
            </div>
        );
    }
}

ReactDOM.render(
    <App />,
    document.getElementById('app')
);