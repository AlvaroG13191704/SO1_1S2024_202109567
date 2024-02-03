import React, { useRef } from 'react';
import 'tailwindcss/tailwind.css';
import Webcam from "react-webcam";

interface Payload {
  imageBase64: string;
  createdAt: string;
}

const App = () => {
  const webcamRef = useRef(null);
  const [capturedImage, setCapturedImage] = React.useState<string | null>(null);
  const [showModal, setShowModal] = React.useState(false);

  const capture = React.useCallback(
    () => {
      if (webcamRef.current) {
        const imageSrc = (webcamRef.current as Webcam).getScreenshot();
        setCapturedImage(imageSrc);
      }
    },
    [webcamRef]
  );


  const sendPicture = async () => {


    if (!capturedImage) {
      return;
    }

    const payload: Payload = {
      imageBase64: capturedImage || "", // Update the type to allow for a string or null value
      createdAt: new Date().toISOString(),
    };

    const response = await fetch('http://127.0.0.1:8080/upload', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload),
    });

    const data = await response.json();
    console.log('Response from API:', data);

    setShowModal(true);

    setTimeout(() => {
      setShowModal(false);
    }, 1000);

  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <h1 className="mb-4 text-2xl font-bold text-gray-900">Tomate una fotografía</h1>
      <Webcam
        audio={false}
        ref={webcamRef}
        screenshotFormat="image/jpeg"
      />
      {capturedImage && (
        <div className="w-64 h-64 overflow-hidden mt-2">
          <img src={capturedImage} alt="Captured" className="w-full h-auto" />
        </div>
      )}
      <div className="space-x-4 mt-2">
        <button
          className="px-4 py-2 font-bold text-white bg-blue-500 rounded hover:bg-blue-700"
          onClick={capture}
        >
          Tomar foto
        </button>
        <button
          className="px-4 py-2 font-bold text-white bg-green-500 rounded hover:bg-green-700"
          onClick={sendPicture}
        >
          Guardar foto
          {showModal && (
            <div className="fixed inset-0 z-10 flex items-center justify-center">
              <div className="absolute inset-0 bg-black opacity-75"></div>
              <div className="z-10 p-4 bg-white rounded-lg">
                <h2 className="mb-2 text-2xl font-bold text-gray-900">Foto guardada</h2>
                <button
                  className="px-4 py-2 font-bold text-white bg-green-500 rounded hover:bg-green-700"
                  onClick={() => setShowModal(false)}
                >
                  Cerrar
                </button>
              </div>
            </div>
          )}
        </button>
      </div>
    </div>
  );
};

export default App;