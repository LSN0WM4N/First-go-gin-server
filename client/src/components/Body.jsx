import { useState } from "react";

const options = [
    {"value" : "Salon", "luminosity" : 250},
    {"value" : "Dormitorios", "luminosity" : 200},
    {"value" : "Cocinas", "luminosity" : 300},
    {"value" : "Cuartos de baño", "luminosity" : 200},
    {"value" : "Pasillos", "luminosity" : 200},
    {"value" : "Mesa de comedor", "luminosity" : 500},
    {"value" : "Espejo del baño", "luminosity" : 500},
    {"value" : "Mesa de lectura", "luminosity" : 600},
]

const getLumensNeeded = (room) => {
    for (let i in options) {
        if (options[i].value === room) {
            return options[i].luminosity
        }
    }
    return -1   
}

// TODO: Make the buttom unaviable when options aren't setted   
const Body = () => {
    const [xSide, setXSIDE] = useState(0);
    const [ySide, setYSIDE] = useState(0);
    const [lum, setLum] = useState(0);
    const [selectedOption, setSelectedOption] = useState('');
    const [answer, setAnswer] = useState(0);

    const handleSubmit = (event) => {
        event.preventDefault();

        const selected = getLumensNeeded(selectedOption)
        if (selected === -1) {
            alert("Debe seleccionar una habitacion para poder realizar el calculo")
            return
        }

        setAnswer(Math.ceil(xSide * ySide * selected / lum))
    };

    return (
        <div className="Body"> 
            <form onSubmit={handleSubmit}>
                
                <label for="options">Habitacion: </label>
                <select 
                    id="options" 
                    value={selectedOption} 
                    className="formObject" 
                    onChange={(e) => setSelectedOption(e.target.value)}
                >
                    
                    <option value="">--Please choose an option--</option>
                    { options.map( (e) => {
                        return <option key={e.value} value={e.value}>{ e.value }</option>
                    } ) } 

                </select>
                
                <label for="lum">Diga Luminosidad: </label>
                <input 
                    className="formObject"
                    id="lum"
                    type="number"
                    onChange={ (e) => {setLum(e.target.value)} }
                ></input>

                <label for="xSide">Diga largo de la habitacion: </label>
                <input
                    className="formObject" 
                    id="xSide"
                    type="number"
                    onChange={ (e) => {setXSIDE(e.target.value)} }
                ></input>
                
                <label for="ySide">Diga ancho de la habitacion: </label>
                <input 
                    className="formObject"
                    id="ySide"
                    type="number"
                    onChange={ (e) => {setYSIDE(e.target.value)} }
                ></input>
                
                <button type="submit" className="buttom">Enviar</button>
            </form>
        
            <form>
                <h3> Se necesitan { answer } luces al menos </h3> 
            </form>
        </div>
    );
}

export default Body;