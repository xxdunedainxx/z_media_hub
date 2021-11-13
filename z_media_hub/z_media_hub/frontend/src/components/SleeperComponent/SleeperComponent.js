import React, { useState } from 'react';
import Modal from 'react-modal';
import './SleeperComponent.css';

function SleeperComponent() {
  const [sleeperMinutes, setSleeperTimer] = useState(60);

  const handleTimerChange = (event) => setSleeperTimer(event.target.value);


  function handleSubmit() {
    window.backend.Sleep(sleeperMinutes)
    alert(`Sleep scheduled in ${sleeperMinutes} minute(s)!`)
  }

  return (
    <div class="sleeperWrapper">
      <label for="sleeperMinutesInput" class="sleeperMinutesInputLabel">
        Enter time to sleep:
      </label>
      <textarea class="sleeperMinutesInput" id="sleeperMinutesInput" value={sleeperMinutes} onChange={handleTimerChange} />
      <br />
      <button class="userInfoSubmitButton" onClick={handleSubmit}>Get some sleep!</button>
    </div>
  );
  
}

export default SleeperComponent;