class App extends React.Component {
  render() {
      return (<Home />);
  }
}

class Home extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      selected_cloud: 'aws',
      selected_service: 'cloudwatchlogs',
      selected_resource: 'log_groups'
    };
  }

  selectResource = (cloud, service, resource) => {
    console.log("selecting resource " + cloud + " " + service + " " + resource);
    this.setState({
      selected_cloud: cloud,
      selected_service: service,
      selected_resource: resource
    })
  }

  render() {
    return (
      <div>
        <SideBar onClick={this.selectResource}/>
        <ResourceView cloud={this.state.selected_cloud} service={this.state.selected_service} resource={this.state.selected_resource}/>
      </div>
    );
  }
}

class ResourceView extends React.Component {
  constructor(props) {
    super(props);

    const today = new Date();
    this.state = {
      objects: [],
      metadata: {},
      filters: {
        report_date: today.toISOString().substring(0, 10),
      },
      populated: false,
    };
  }

  serverRequest = () => {
    var inventoryURL = `api/v1/inventory/${this.props.cloud}/${this.props.service}/${this.props.resource}?report_date=${this.state.filters.report_date}`;
    var metadataURL = `api/v1/metadata/${this.props.cloud}/${this.props.service}/${this.props.resource}?report_date=${this.state.filters.report_date}`
    if ('report_time' in this.state.filters) {
      inventoryURL += `&time_selection=at&time_selection_reference=${this.state.filters.report_time}`;
    }
    Promise.all([
      $.get(inventoryURL),
      $.get(metadataURL),
    ]).then(([objects, metadata]) => {
      this.setState({
        objects: objects,
        metadata: metadata,
        populated: true,
      });
    })
  }

  componentDidMount() {
    this.serverRequest();
  }

  componentDidUpdate(prevProps, prevState) {
    if (this.props.cloud !== prevProps.cloud || this.props.service !== prevProps.service || this.props.resource !== prevProps.resource || this.state.filters !== prevState.filters) {
      console.log("making new request")
      this.serverRequest();
    }
  }

  render() {
    if (!this.state.populated) {
      console.log("not populated");
      return (<div/>)
    }
    console.log(this.state.metadata)
    return (
      <div class="main">
        <br />
        <h2>Cloud Inventory</h2>
        <p>{this.props.cloud} {this.props.service} {this.props.resource}</p>
        <input type="date" value={this.state.filters.report_date} onChange={(e) => {
          this.setState({filters: {...this.state.filters, report_date: e.target.value}})
        }} />
        <div class="dropdown">
          <button class="btn btn-primary dropdown-toggle" type="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Date/Time Selection
            <span class="caret"></span>
          </button>
          <ul class="dropdown-menu">
            {this.state.metadata.datetimes.map((datetime) => {
              var state = ""
              if ('report_time' in this.state.filters && this.state.filters.report_time === datetime) {
                  state = "active"
              }
              return (
                <li class={state}>
                  <a onClick={() => {
                    console.log("clicked " + datetime);
                    this.setState({filters: {...this.state.filters, report_time: datetime}})
                  }}>{datetime}</a>
                </li>
              )
            })}
          </ul>
        </div>
        <div class="container">
          <div class="panel-group">
            {this.state.objects.map((object, i) => {
              return <InventoryObject object={object} index={i} id_field={this.state.metadata.id_field}/>;
            })}
          </div>
        </div>
      </div>
    )
  }
}

class InventoryObject extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      open: false
    };
  }

  render() {
    return (
      <div class="panel panel-default">
        <div class="panel-heading">
          <h4 class="panel-title">
            <a data-toggle="collapse" href={"#"+this.props.index}>{this.props.object[this.props.id_field]}</a>
          </h4>
        </div>
        <div id={this.props.index} class="panel-collapse collapse">
          <div class="panel-body"><pre>{JSON.stringify(this.props.object, null, 2)}</pre></div>
        </div>
      </div>
    )
  }
}

class SideBar extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      clouds: ["aws"]
    }
  }

  render() {
    return (
      <div class="sidenav">
        <div class="panel-group">
          {this.state.clouds.map((cloud) =>{
            return <CloudTab cloud={cloud} onClick={this.props.onClick}/>;
          })}
        </div>
      </div>
    )
  }
}

class CloudTab extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      services: []
    }
  }

  fetchServices = () => {
    $.get(`api/v1/metadata/${this.props.cloud}`, res => {
      this.setState({
        services: res.services
      });
    });
  }
8
  componentDidMount() {
    this.fetchServices();
  }

  render() {
    return (
      <div class="panel panel-default">
        <div class="panel-heading">
          <h4 class="panel-title">
            <a data-toggle="collapse" href={"#"+this.props.cloud}>{this.props.cloud}</a>
          </h4>
        </div>
        <div id={this.props.cloud} class="panel-collapse collapse">
            <div class="panel-body">
              {this.state.services.map((service) => {
                return <ServiceTab cloud={this.props.cloud} service={service} onClick={this.props.onClick}/>;
              })}
            </div>
          </div>
      </div>
    )
   
  }
}

class ServiceTab extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      resources: []
    }
  }

  fetchResources = () => {
    $.get(`api/v1/metadata/${this.props.cloud}/${this.props.service}`, res => {
      this.setState({
        resources: res.resources
      });
    });
  }

  componentDidMount() {
    this.fetchResources();
  }

  render() {
    return (
      <div class="panel panel-default">
        <div class="panel-heading">
          <h4 class="panel-title">
            <a data-toggle="collapse" href={"#"+this.props.cloud+"_"+this.props.service}>{this.props.service}</a>
          </h4>
        </div>
        <div id={this.props.cloud+"_"+this.props.service} class="panel-collapse collapse">
          <div class="panel-body">
            <ul class="list-group">
              {this.state.resources.map((resource) => {
                return (
                  <li class="list-group-item">
                    <a href="#" onClick={() => this.props.onClick(this.props.cloud, this.props.service, resource)}>{resource}</a>
                  </li>
                )
              })}
            </ul>
          </div>
        </div>
      </div>
    )
  }
}

ReactDOM.render(<App />, document.getElementById('app'));
