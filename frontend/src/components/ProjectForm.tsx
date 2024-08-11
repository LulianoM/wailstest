import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { GenerateProject } from "../../wailsjs/go/main/App";

interface ProjectFormProps {
  title: string;
}

const ProjectForm: React.FC<ProjectFormProps> = ({ title }) => {
  const [projectName, setProjectName] = useState('');
  const [teamName, setTeamName] = useState('');
  const [resultText, setResultText] = useState('');

  const navigate = useNavigate();

  const updateResultText = (result: string) => setResultText(result);

  const handleGenerateProject = () => {
    GenerateProject(projectName, teamName).then(setResultText);
  };

  const goToHome = () => {
    navigate('/');
  };

  return (
    <div className="container">
      <h1 className="title">{title}</h1>
      <div className="form-group">
        <label htmlFor="projectName">Nome do projeto</label>
        <input
          id="projectName"
          type="text"
          value={projectName}
          onChange={(e) => setProjectName(e.target.value)}
        />
      </div>
      <div className="form-group">
        <label htmlFor="teamName">Nome do Time do Projeto</label>
        <input
          id="teamName"
          type="text"
          value={teamName}
          onChange={(e) => setTeamName(e.target.value)}
        />
      </div>
      <button className="btn" onClick={handleGenerateProject}>
        Processar
      </button>
      {resultText && (
        <div id="result" className="result">
          {resultText}
        </div>
      )}
      <button className="btn back-btn" onClick={goToHome}>
        Voltar para Home
      </button>
    </div>
  );
};

export default ProjectForm;
