import { useEffect } from 'react'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { SignupFormInterface } from '@/features/auth/types/types'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { signup } from '@/features/auth/api/signup'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { AlertCircleIcon } from 'lucide-react'


const signupSchema = z.object({
  login: z.string().min(2, "Identifiant trop court"),
  password: z.string().min(6, "Mot de passe trop court"),
});

// type SignupFormValues = z.infer<typeof signupSchema>

export function SignupForm() {
    const { register, setFocus, handleSubmit, setError, formState: { errors } } = useForm<SignupFormInterface>({
        resolver: zodResolver(signupSchema),
    })

    const onSubmit = async (data: SignupFormInterface) => {
        await signup(data)
            .catch((error) => {
                setError('root', { message: error.message })
            })
    }

    useEffect(() => {
        setFocus('login')
    }, [setFocus])

    return (
        <form className="space-y-4 max-w-sm mx-auto mt-8" onSubmit={handleSubmit(onSubmit)}>
            {errors.root && (
                <div className="grid w-full max-w-xl items-start gap-4">
                    
                    <Alert variant="destructive">
                    <AlertCircleIcon />
                    <AlertDescription>
                        <p>{errors.root.message}</p>
                    </AlertDescription>
                    </Alert>
                </div>
            )}
            <div>
                <label htmlFor="login" className="block mb-1">
                    Identifiant
                </label>
                <Input
                    id="login"
                    type="text"
                    placeholder="Entrez votre identifiant"
                    {...register("login")}
                    required
                />
            </div>
            <div>
                <label htmlFor="password" className="block mb-1">
                    Mot de passe
                </label>
                <Input
                    id="password"
                    type="password"
                    placeholder="Votre mot de passe"
                    {...register("password")}
                    required
                />
            </div>
            <Button type="submit" className="w-full">
                S’inscrire
            </Button>
        </form>
    )
}