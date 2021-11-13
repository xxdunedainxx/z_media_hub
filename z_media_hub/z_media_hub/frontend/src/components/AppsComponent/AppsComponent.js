import React, { useState } from 'react';
import Modal from 'react-modal';
import netflixLogo from './assets/netflix.jpeg'
import amazonLogo from './assets/amazon.png'
import huluLogo from './assets/hulu.jpeg'
import disneyLogo from './assets/disney.jpeg'
import hboLogo from './assets/hbo.jpeg'
import youtubeLogo from './assets/youtube.gif'

import './AppsComponent.css';

function AppsComponent() {

  function handleNetflix() {
    window.backend.Netflix()
  }
  
  function handleHulu() {
    window.backend.Hulu()
  }
  
  function handleAmazon() {
    window.backend.Amazon()
  }

  function handleHBO() {
    window.backend.HBO()
  }

  function handleYouTube() {
    window.backend.YouTube()
  }

  function handleDisney() {
    window.backend.Disney()
  }

  return (
    <div class="appsComponent">
    <h3 class="appsComponentHeader">Z's Media Hub ;)</h3>
      <img src={netflixLogo} alt="netflix?" onClick={handleNetflix} className="appClickable" />
      <img src={huluLogo} alt="hulu?" onClick={handleHulu} className="appClickable" />
      <img src={amazonLogo} alt="amazon?" onClick={handleAmazon} className="appClickable" />
      <img src={hboLogo} alt="hbo?" onClick={handleHBO} className="appClickable" />
      <img src={youtubeLogo} alt="youtube?" onClick={handleYouTube} className="appClickable" />
      <img src={disneyLogo} alt="disney?" onClick={handleDisney} className="appClickable" />
    </div>
  );
  
}

export default AppsComponent;