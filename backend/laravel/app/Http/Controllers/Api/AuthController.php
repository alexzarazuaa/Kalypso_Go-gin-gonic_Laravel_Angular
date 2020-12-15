<?php

namespace App\Http\Controllers\Api;

use Auth;
use App\User;
use App\Http\Requests\Api\LoginUser;
use App\Http\Requests\Api\RegisterUser;
use App\RealWorld\Transformers\UserTransformer;

class AuthController extends ApiController
{
    /**
     * AuthController constructor.
     *
     * @param UserTransformer $transformer
     */
    public function __construct(UserTransformer $transformer)
    {
        $this->transformer = $transformer;
    }

    /**
     * Login user and return the user if successful.
     *
     * @param LoginUser $request
     * @return \Illuminate\Http\JsonResponse
     */
    public function login(LoginUser $request)
    {
        $credentials = $request->only('user.email', 'user.password');
        $credentials = $credentials['user'];


        // echo ("oooooooooooooo");


        if (!Auth::once($credentials)) {

            return $this->respondFailedLogin();
        }

        // echo ("HOLAGHOLAHOLAGOLAAASASA");

        return $this->respondWithTransformer(auth()->user());
    }


    public function login_admin_go(LoginUser $request)
    {
        $credentials = $request->only('user.email');
        $credentials = $credentials['user'];

        $user_id = User::all()->where('email', $credentials['email'])->first();

        // print_r($user_id['password']);

        $password=$user_id['password'];

        print_r($password);

    }

    /**
     * Register a new user and return the user if successful.
     *
     * @param RegisterUser $request
     * @return \Illuminate\Http\JsonResponse
     */
    public function register(RegisterUser $request)
    {
        $user = User::create([
            'username' => $request->input('user.username'),
            'email' => $request->input('user.email'),
            'password' => bcrypt($request->input('user.password')),
        ]);

        return $this->respondWithTransformer($user);
    }
}
