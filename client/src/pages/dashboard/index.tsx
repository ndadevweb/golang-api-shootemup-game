import { Link } from 'react-router'
import { ROUTES } from '@/routes';

export default function DashboardPage() {
    const user = null
    return (
        <>
            <h1>Accueil utilisateur identifié</h1>
            <p>Bienvenue <strong>{ user || 'Utilisateur inconnu pour le moment...' }</strong></p>
        </>
    );
}
