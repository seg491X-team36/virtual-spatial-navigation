import { Outlet, Link } from "react-router-dom";

const Layout = () => {
  return (
    <>
      <nav>
        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
          <li>
            <Link to="/experiments">View Experiments</Link>
          </li>
          <li>
            <Link to="/participants">View Participants</Link>
          </li>
          <li>
            <Link to="/rejected-sign-ups">View Rejected Sign Ups</Link>
          </li>
          <li>
            <Link to="/sign-ups">View Sign Ups</Link>
          </li>
        </ul>
      </nav>

      <Outlet />
    </>
  )
};

export default Layout;