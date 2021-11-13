import React from 'react';
import logo from './logo.png';
import './App.css';
import SleeperComponent from './components/SleeperComponent/SleeperComponent';
import AppsComponent from './components/AppsComponent/AppsComponent';
import VersionComponent from './components/VersionComponent/VersionComponent';
function App() {
  return (
    <div id="app" className="App">
      <header className="App-header">
        <VersionComponent />
        <AppsComponent />
        <SleeperComponent />
      </header>
    </div>
  );
}

export default App;
