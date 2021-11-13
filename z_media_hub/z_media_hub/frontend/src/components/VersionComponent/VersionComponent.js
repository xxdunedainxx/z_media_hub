import React, { useState } from 'react';
import Modal from 'react-modal';
import './VersionComponent.css';

function VersionComponent() {
  const VERSION = '0.0.1'

  return (
    <div class="versionWrapper">
     Version {VERSION}
    </div>
  );
  
}

export default VersionComponent;