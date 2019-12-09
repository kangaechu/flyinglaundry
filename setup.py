import os, sys
from setuptools import setup, find_packages

def read_requirements():
    """Parse requirements from requirements.txt."""
    reqs_path = os.path.join('.', 'requirements.txt')
    with open(reqs_path, 'r') as f:
        requirements = [line.rstrip() for line in f]
    return requirements

setup(
    name='flying_laundry',
    version='0.0.1',
    description='Check whether your laundries can fly.',
    long_description=Readme.md,
    author='kangaechu',
    author_email='kangae2@gmail.com',
    install_requires=read_requirements(),
    url='https://github.com/kangaechu/flyinglaundry',
    license='mit',
    packages=find_packages(exclude=('tests', 'docs'))
)