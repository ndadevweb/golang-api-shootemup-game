import { Link } from 'react-router'
import { ROUTES } from '@/routes';

export default function HomePage() {
  return (
    <>
        <h1>Accueil</h1>
        <ul className="mt-5">
            <li><Link to={ROUTES.signup}>M'inscrire</Link></li>
            <li><Link to={ROUTES.signin}>Me connecter</Link></li>
        </ul>
    </>
  );
}
