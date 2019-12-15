import os, sys
from setuptools import setup, find_packages

sys.path.append('./flyinglaundry')
sys.path.append('./test')

def read_requirements():
    """Parse requirements from requirements.txt."""
    reqs_path = os.path.join('.', 'requirements.txt')
    with open(reqs_path, 'r') as f:
        requirements = [line.rstrip() for line in f]
    return requirements

setup(
    name='flyinglaundry',
    version='0.0.1',
    description='Check whether your laundries can fly.',
    long_description='README.md',
    author='kangaechu',
    author_email='kangae2@gmail.com',
    install_requires=read_requirements(),
    url='https://github.com/kangaechu/flyinglaundry',
    license='mit',
    test_suite = 'test',
    packages=find_packages(exclude=('tests', 'docs'))
)