import React from 'react'
import { connect } from 'react-redux'
import {replaceRoute} from '../../actions/misc';
/**
 * Wraps components so they cannot be accessed if user is not authenticated.
 *
 * @param Component
 * @returns {*}
 */
export default function requireAuthentication(Component) {

  class AuthenticatedComponent extends React.Component {
    componentWillMount() {
      this.checkAuth(this.props.auth)
    }
    componentWillReceiveProps(nextProps) {
      this.checkAuth(nextProps.auth)
    }
    checkAuth(auth) {
      if (!auth.isAuthenticated) {
        this.props.dispatch(replaceRoute('/about'))
      }
    }
    render() {
      return (
        <div>
          {this.props.auth.isAuthenticated === true
            ? <Component {...this.props} />
            : null
          }
        </div>
      )
    }
  }

  function mapStateToProps(state) {
    return {
      auth: state.auth
    }
  }

  return connect(mapStateToProps)(AuthenticatedComponent)
}
