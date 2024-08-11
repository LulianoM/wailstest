import React from 'react';
import { Link } from 'react-router-dom';

const Home: React.FC = () => {
  return (
    <div className="container">
      <h1 className="title">Admin Tools - BFF</h1>
      <div className="button-group">
        <Link to="/create-bff">
          <button className="btn">Create New BFF</button>
        </Link>
        <Link to="/create-mf">
          <button className="btn">Create New MF</button>
        </Link>
      </div>
    </div>
  );
};

export default Home;
